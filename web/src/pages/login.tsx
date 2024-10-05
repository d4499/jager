import { createForm, SubmitHandler, valiForm } from "@modular-forms/solid";
import { Match, Switch } from "solid-js";
import * as v from "valibot";

const LoginSchema = v.object({
	email: v.pipe(
		v.string(),
		v.nonEmpty("Please enter your email."),
		v.email("Please enter a valid email"),
		v.maxLength(320),
	),
});

type LoginForm = v.InferInput<typeof LoginSchema>;

async function requestMagicLink(email: string) {
	const res = await fetch("http://localhost:8080/api/auth/magic", {
		method: "POST",
		credentials: "include",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ email }),
	});

	return res.json();
}

export function Login() {
	const [loginForm, { Form, Field }] = createForm<LoginForm>({
		validate: valiForm(LoginSchema),
	});

	const handleSubmit: SubmitHandler<LoginForm> = (values) => {
		requestMagicLink(values.email);
	};

	return (
		<div>
			<Switch
				fallback={
					<>
						<h1>Login</h1>
						<Form onSubmit={handleSubmit}>
							<Field name="email">
								{(field, props) => (
									<>
										<input {...props} type="email" />
										{field.error && <div>{field.error}</div>}
									</>
								)}
							</Field>
							<button type="submit">Continue with Email</button>
						</Form>
					</>
				}
			>
				<Match when={loginForm.submitted}>
					<p>Please check your email to continue</p>
				</Match>
			</Switch>
		</div>
	);
}
