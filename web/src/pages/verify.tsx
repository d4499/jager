import { useNavigate, useSearchParams } from "@solidjs/router";
import { createEffect } from "solid-js";
import { useAuth } from "../providers/auth";

export function Verify() {
	const { verifyMagicLink, isAuthenticated } = useAuth();
	const [searchParams] = useSearchParams();
	const nav = useNavigate();

	createEffect(async () => {
		const token = searchParams.token;
		if (token) {
			const verified = await verifyMagicLink(token);
			if (!verified) {
				nav("/login", { replace: true });
			}
			nav("/protected", { replace: true });
		}
	});

	return (
		<div>
			<p>token {searchParams.token}</p>
			<p>authenticated: {isAuthenticated()}</p>
		</div>
	);
}
