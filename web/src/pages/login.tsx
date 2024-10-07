import { createForm, SubmitHandler, valiForm } from "@modular-forms/solid";
import { Match, Switch } from "solid-js";
import * as v from "valibot";
import { useAuth } from "../providers/auth";

const LoginSchema = v.object({
	email: v.pipe(
		v.string(),
		v.nonEmpty("Please enter your email."),
		v.email("Please enter a valid email"),
		v.maxLength(320),
	),
});

type LoginForm = v.InferInput<typeof LoginSchema>;

export function Login() {
	const { sendMagicLink } = useAuth();
	const [loginForm, { Form, Field }] = createForm<LoginForm>({
		validate: valiForm(LoginSchema),
	});

	const handleSubmit: SubmitHandler<LoginForm> = (values) => {
		sendMagicLink(values.email);
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
