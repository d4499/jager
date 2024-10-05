import { useSearchParams } from "@solidjs/router";
import { createEffect } from "solid-js";
import { useAuth } from "../providers/auth";

export function Verify() {
	const { verifyMagicLink, isAuthenticated } = useAuth();
	const [searchParams] = useSearchParams();

	createEffect(() => {
		const token = searchParams.token;
		if (token) {
			verifyMagicLink(token);
		}
	});

	return (
		<div>
			<p>token {searchParams.token}</p>
			<p>authenticated: {isAuthenticated()}</p>
		</div>
	);
}
