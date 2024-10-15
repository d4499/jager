import { createForm, SubmitHandler, valiForm } from "@modular-forms/solid";
import * as v from "valibot";
import { useAuth } from "../providers/auth";
import { Button } from "../components/ui/button";
import {
	TextField,
	TextFieldLabel,
	TextFieldInput,
} from "../components/ui/text-field";
import { createSignal, Match, Switch } from "solid-js";

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
	const [sent, setSent] = createSignal(false);
	const { sendMagicLink } = useAuth();
	const [_, { Form, Field }] = createForm<LoginForm>({
		validate: valiForm(LoginSchema),
	});

	const handleSubmit: SubmitHandler<LoginForm> = (values) => {
		sendMagicLink(values.email);
		setSent(true);
	};

	return (
		<div>
			<h1>Login</h1>
			<Switch
				fallback={
					<>
						<Form onSubmit={handleSubmit}>
							<Field name="email">
								{(field, props) => (
									<TextField class="grid w-full max-w-sm items-center gap-2">
										<TextFieldLabel for="email">Email</TextFieldLabel>
										<TextFieldInput {...props} type="email" />
										{field.error && (
											<div class="text-red-500">{field.error}</div>
										)}
									</TextField>
								)}
							</Field>
							<Button type="submit">Continue with Email</Button>
						</Form>
					</>
				}
			>
				<Match when={sent()}>
					<p>Please check your email</p>
				</Match>
			</Switch>
		</div>
	);
}
