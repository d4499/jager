import { Route, Router } from "@solidjs/router";
import { Root } from "./pages/root";
import { Login } from "./pages/login";

function App() {
	return (
		<Router>
			<Route path="/" component={Root} />
			<Route path="/login" component={Login} />
		</Router>
	);
}

export default App;
