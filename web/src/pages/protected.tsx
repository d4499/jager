import { useAuth } from "../providers/auth";

export function Protected() {
	const { session } = useAuth();
	return (
		<div>
			<h1>Protected</h1>
			<h1>{session()?.userId}</h1>
		</div>
	);
}
