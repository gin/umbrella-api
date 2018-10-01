class ProposalHistory extends React.Component {
  render() {
    <div>
      <h3>Proposal History</h3>
      <ul>
        {this.props.proposals.map(p => <li>{p}</li>)}
      </ul>
    </div>
  }
}
