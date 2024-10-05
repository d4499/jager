import { Route, Router } from "@solidjs/router";
import { Root } from "./pages/root";
import { Login } from "./pages/login";
import { Verify } from "./pages/verify";
import { AuthProvider } from "./providers/auth";
import { Protected } from "./pages/protected";
import { AuthGuard } from "./components/auth-guard";

function App() {
	return (
		<AuthProvider>
			<Router>
				<Route path="/" component={Root} />
				<Route path="/login" component={Login} />
				<Route path="/verify" component={Verify} />
				<Route component={AuthGuard}>
					<Route path="/protected" component={Protected} />
				</Route>
			</Router>
		</AuthProvider>
	);
}

export default App;
