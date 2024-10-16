import { createForm, SubmitHandler, valiForm } from "@modular-forms/solid";
import * as v from "valibot";

async function createJobApplication(
	title: string,
	company: string,
	applied_date: Date,
) {
	const res = await fetch("http://localhost:8080/api/jobapplications/", {
		method: "POST",
		credentials: "include",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			title: title,
			company: company,
			applied_date: applied_date,
		}),
	});
	const data = await res.json();

	return data;
}

const CreateApplicationSchema = v.object({
	title: v.pipe(v.string(), v.minLength(2), v.maxLength(32)),
	company: v.pipe(v.string(), v.minLength(2), v.maxLength(32)),
	applied_date: v.pipe(v.date(), v.maxValue(new Date())),
});

type CreateApplication = v.InferInput<typeof CreateApplicationSchema>;

export function CreateJobApplication() {
	const [_, { Form, Field }] = createForm<CreateApplication>({
		validate: valiForm(CreateApplicationSchema),
	});

	const handleSubmit: SubmitHandler<CreateApplication> = (values) => {
		console.log("submitted");
		createJobApplication(values.title, values.company, values.applied_date);
	};

	return (
		<Form onSubmit={handleSubmit}>
			<Field name="title">
				{(field, props) => (
					<>
						<input {...props} type="text" />
						{field.error && <div>{field.error}</div>}
					</>
				)}
			</Field>
			<Field name="company">
				{(field, props) => (
					<>
						<input {...props} type="text" />
						{field.error && <div>{field.error}</div>}
					</>
				)}
			</Field>
			<Field name="applied_date" type="Date">
				{(field, props) => (
					<>
						<input {...props} type="date" />
						{field.error && <div>{field.error}</div>}
					</>
				)}
			</Field>
			<button type="submit">Create</button>
		</Form>
	);
}
