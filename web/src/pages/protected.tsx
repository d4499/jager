import { createResource, For, Match, Switch } from "solid-js";
import { CreateJobApplication } from "../components/create-job-application";

type JobApplication = {
	id: string;
	title: string;
	company: string;
	user_id: string;
	applied_date: string;
	created_at: Date;
	updated_at: Date;
};

async function fetchJobApplications(): Promise<JobApplication[]> {
	const res = await fetch("http://localhost:8080/api/jobapplications", {
		credentials: "include",
	});
	const data = await res.json();

	return data;
}

export function Protected() {
	const [jobApplications] = createResource(fetchJobApplications);

	return (
		<div>
			<h1>Protected</h1>
			<Switch>
				<Match when={jobApplications.loading}>
					<p>Loading...</p>
				</Match>
				<Match when={jobApplications.error}>
					<p>Error</p>
				</Match>
				<Match when={jobApplications()}>
					<For each={jobApplications()}>
						{(application) => (
							<div>
								<p>{application.title}</p>
								<p>{application.company}</p>
							</div>
						)}
					</For>
				</Match>
			</Switch>
			<CreateJobApplication />
		</div>
	);
}
