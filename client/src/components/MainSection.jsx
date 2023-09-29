import { useState, useEffect } from "react";
import Card from "./UI/Card";

const MainSection = () => {
	const [issues, setIssues] = useState([]);
	const language = "javascript";

	useEffect(() => {
		fetch(`http://localhost:10000/api/issues/${language}`)
			.then((response) => {
				if (!response.ok) {
					throw new Error(`Network response was not ok: ${response.status}`);
				}
				return response.json();
			})
			.then((data) => {
				let results = data;
				setIssues(results);
				localStorage.setItem("Issues", JSON.stringify(results));
			})
			.catch((error) => {
				console.error("There was a problem with the fetch operation:", error);
			});

		console.log("Issues fetched successfully", issues);
	}, [language]);

	return (
		<main className='h-full pb-6 pt-24 px-6 w-full sm:max-w-[65dvw] fixed right-0 bg-primary min-h-screen overflow-y-scroll'>
			<h1 className='text-white uppercase tracking-wider text-center mb-6'>Open issues</h1>

			<div className='flex flex-col gap-6 w-full sm:w-4/5 mx-auto'>
				{issues?.length > 0 ? (
					issues.map((issue) => <Card key={issue?.html_url} {...issue} />)
				) : (
					<p className='text-white'>No issues found</p>
				)}
			</div>
		</main>
	);
};

export default MainSection;
