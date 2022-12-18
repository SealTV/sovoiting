// SPDX-Lincense-Identifier: GPL-3.0
// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity >= 0.8.17;

// @title Voting with delegation.
contract Ballot {
    // This declares a new complex type which will
    // be used for variables later.
    // It will represent a single voter.
    struct Voter {
        uint weight; // weight is accumulated be delegation
        bool voted; // if true, that person already voted
        address delegate; // person delegated to
        uint vote; // index of the voted proposal
    }

    // This is stuct for a single proposal.
    struct Proposal {
        bytes32 name; // short name (up to 32 bytes)
        uint voteCount; // number of accumulated votes
    }

    address public chairperson;

    // This declares a state variable that
    // soters as 'Voter' stuct for each possible address.
    mapping(address => Voter) public voters;

    // A dynamically-sized array of `Proposal` structs.
    Proposal[] public proposals;

    
    enum VoteStates { Created, Active, Ended }

    VoteStates state;
    VoteStates constant defaultState = VoteStates.Created;

    /// Create new ballot to chose one of 'proposalNames'.
    constructor(bytes32[] memory proposalNames) {
        chairperson = msg.sender;
        voters[chairperson].weight = 1;

        // For each of provided proposal names,
        // create new proposal object and add it
        // to the end of array.
        for (uint i = 0; i < proposalNames.length; i++) {
            proposals.push(Proposal({name: proposalNames[i], voteCount: 0}));
        }

        state = defaultState;
    }

    function startVoting() external {
        require(
            msg.sender == chairperson,
            "Only chairperson start voting."
        );

        state = VoteStates.Active;
    }

    function endVoting() external {
        require(
            msg.sender == chairperson,
            "Only chairperson end voting."
        );

        state = VoteStates.Ended;
    }

    function isVotingActive() public view returns (bool isActive) {
        isActive = state == VoteStates.Active;
    }

    // Give 'voter' right to vite on this ballot.
    // May only be called by 'chairperson'.
    function giveRightToVote(address voter) external {
        // If the first of 'require' evaluates
        // to `false`, execution terminated and all
        // changes to the state and to Ether balances
        // are reverted.
        // This used to consume all gas in old EVM version, but
        // not anymore.
        // It is often good idea to use 'require` to check if
        // function are colled correctly.
        // As a second argument, you can also provide an
        // expanation about what went wrong.
        require(
            msg.sender == chairperson,
            "Only chairperson can give right to vote."
        );

        require(!voters[voter].voted, "The voter already voted.");

        require(voters[voter].weight == 0);

        voters[voter].weight = 1;
    }

    /// Delegate your vote to the voter `to`.
    function delegate(address to) external {
        // assign reference
        Voter storage sender = voters[msg.sender];
        require(sender.weight != 0, "You have no right to vote.");
        require(!sender.voted, "You already voted.");
    
        require(to != msg.sender, "Self-delegation is disallowed.");

        // Forward the delegation as long as
        // `to` also delegated.
        // In general, such loops are very dangerous,
        // because if they run too long, they might
        // need more gas than is available in a block.
        // In this case, the delegation will not be executed,
        // but in other situations, such loops might
        // cause a contract to get "stuck" completely.
        while (voters[to].delegate != address(0)) {
            to = voters[to].delegate;

            require(to != msg.sender, "Found loop in delegation.");
        }

        Voter storage delegate_ = voters[to];

        // Voters cannot delegate to accounts that cannot vote.
        require(delegate_.weight >= 1); 

        // Since `sender` is a reference, this
        // modifies `voters[msg.sender]`.
        sender.voted = true;
        sender.delegate = to;

        if (delegate_.voted) {
            // If the delegate already voted,
            // directly add to the number of votes
            proposals[delegate_.vote].voteCount = sender.weight;
        } else {
            // If the delegate did not vote yet,
            // add to her weight.
            delegate_.weight += sender.weight;
        }
    }

    /// Give your vote (incuding votes delegated to you)
    /// to proposal proposals[proposal].name
    function vote(uint proposal) external {
        Voter storage sender = voters[msg.sender];
        require(sender.weight != 0, "Has no right to vote.");
        require(!sender.voted, "Already voted.");
        require(isVotingActive(), "Voting should be Active");

        sender.voted = true;
        sender.vote = proposal;

        // If `proposal` is out of the range of the array,
        // this will throw automatically and revert all
        // changes.
        proposals[proposal].voteCount += sender.weight;
    }

    /// @dev Computes the winning proposal taking all
    /// previous votes into account.
    function winningProposal() public view returns (uint winningProposal_) {
        uint winningVoteCount = 0;
        for (uint p = 0; p < proposals.length; p++) {
            if (proposals[p].voteCount > winningVoteCount) {
                winningVoteCount = proposals[p].voteCount;
                winningProposal_ = p;
            }
        }
    }

    // Calls winningProposal() function to get the index
    // of the winner contained in the proposals array and then
    // returns the name of the winner
    function winnerNane() external view returns (bytes32 name) {
        name = proposals[winningProposal()].name;
    }
}
