function App() {
    return (
        <div id="main">
            <h1>Hello Profile!</h1>
            <ProfileList />
        </div>
    );
}


class ProfileList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            profiles: [],
            is_active: false,
            current_profile: {
                Name: "",
                Age: 0,
                Gender: "",
                FavoriteFoods: [{ "Food": "ラーメン" }, { "Food": "りんご" }, { "Food": "レモン" }],
            }
        };
    }

    componentDidMount() {
        axios.get('./Profile')
            .then(res => {
                console.log(res.data);
                this.setState({ profiles: this.state.profiles.concat(res.data) });
            })
            .catch((err) => {
                console.log(err);
            });
    }

    get_profile(name) {
        axios.get('./Profile/' + name)
            .then(res => {
                console.log(res.data);
                this.setState({ current_profile: res.data });
                this.setState({ is_active: true });
            })
            .catch((err) => {
                console.log(err);
            });
    }

    render() {
        return (
            <div id="profile">
                <ul>
                    {this.state.profiles
                        .map(profile => (
                            <li key={profile.Name}>
                                <a href="#" onClick={() => this.get_profile(profile.Name)}>{profile.Name}</a>
                            </li>))
                    }
                </ul>
                <hr />
                {this.state.is_active && <Profile current={this.state.current_profile} />}
            </div>
        );
    }
}

class Profile extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        let foods = [];
        for (const food of this.props.current.FavoriteFoods) {
            foods.push(food['Food']);
        }

        return (
            <ul>
                <li>Name: {this.props.current.Name}</li>
                <li>Age: {this.props.current.Age}</li>
                <li>Gender: {this.props.current.Gender}</li>
                <li>Favorite Foods: [{foods.join(" ")}]</li>
            </ul>
        );
    }
}

const target = document.querySelector('#app');
ReactDOM.render(<App />, target);