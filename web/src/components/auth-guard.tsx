import { ParentProps, Show } from "solid-js";
import { useAuth } from "../providers/auth";
import { useNavigate } from "@solidjs/router";

export function AuthGuard(props: ParentProps) {
	const { isAuthenticated, getUser } = useAuth();
	const nav = useNavigate();

	const checkAuth = async () => {
		if (!isAuthenticated()) {
			const user = await getUser();
			if (!user) {
				nav("/login", { replace: true });
			}
		}
	};

	checkAuth();

	return (
		<Show when={isAuthenticated} fallback={<p>Checking authentication</p>}>
			<div>{props.children}</div>
		</Show>
	);
}
