import {
	Accessor,
	createContext,
	createEffect,
	createSignal,
	JSX,
	useContext,
} from "solid-js";

type Session = {
	userId: string;
};

type AuthContextType = {
	isAuthenticated: Accessor<boolean>;
	sendMagicLink: (email: string) => Promise<boolean>;
	verifyMagicLink: (token: string) => Promise<boolean>;
	session: Accessor<Session | null>;
	getUser: () => Promise<Session | null>;
};

const AuthContext = createContext<AuthContextType>();

export function AuthProvider(props: { children: JSX.Element }) {
	const [isAuthenticated, setIsAuthenticated] = createSignal(false);
	const [session, setSession] = createSignal<Session | null>(null);

	async function sendMagicLink(email: string) {
		const res = await fetch("http://localhost:8080/api/auth/magic", {
			method: "POST",
			credentials: "include",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ email }),
		});
		if (res.ok) {
			return true;
		}
		return false;
	}

	async function verifyMagicLink(token: string) {
		try {
			const res = await fetch("http://localhost:8080/api/auth/magic/verify", {
				method: "POST",
				credentials: "include",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ token }),
			});

			return res.ok;
		} catch (e) {
			console.error("Error verifying magic link", e);
			return false;
		}
	}

	async function getUser() {
		try {
			const res = await fetch("http://localhost:8080/api/auth/me", {
				credentials: "include",
			});

			if (res.ok) {
				const sessionData: Session = await res.json();
				setSession(sessionData);
				setIsAuthenticated(true);
				return sessionData;
			} else {
				setSession(null);
				setIsAuthenticated(false);
			}
		} catch (e) {
			setSession(null);
			setIsAuthenticated(false);
		}
		return null;
	}

	const checkAuth = async () => {
		await getUser();
	};

	createEffect(() => {
		checkAuth();
	});

	const values: AuthContextType = {
		isAuthenticated,
		sendMagicLink,
		verifyMagicLink,
		session,
		getUser,
	};

	return (
		<AuthContext.Provider value={values}>{props.children}</AuthContext.Provider>
	);
}

export function useAuth() {
	const context = useContext(AuthContext);
	if (!context) {
		throw new Error("useAuth must be used within a AuthProvider");
	}
	return context;
}
