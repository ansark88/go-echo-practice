function App() {
    return (
        <div>
            <h1>profile</h1>
            <ProfileList />
        </div>
    );
}

class ProfileList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            profiles: []
        };
    }

    componentDidMount() {
        axios
            .get('./Profile/Alice')
            .then(res => {
                console.log(res.data);
                this.setState({ profiles: this.state.profiles.concat(res.data) });
            })
            .catch((err) => {
                console.log(err);
            });
    }

    render() {
        const profile = this.state.profiles;

        return (
            <ul>
                { this.state.profiles.map(profile => <li key={profile.Name}>{profile.Name}</li>)}
            </ul>
        )
    }
}

const target = document.querySelector('#app');
ReactDOM.render(<App />, target);