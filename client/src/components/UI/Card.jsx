import Badge from "./Badge";

const Card = ({
	name,
	description,
	open_issues_count,
	html_url,
	pushed_at,
	stargazers_count,
	language,
}) => {
	return (
		<div className='bg-primary border-2 border-gray-500 hover:bg-sideBg rounded-md p-4 shadow-md shadow-gray-700 text-white flex gap-2 justify-between'>
			<div className='max-w-[200px] sm:max-w-[400px]'>
				<a
					href={html_url}
					target='_blank'
					className='text-xl text-accentTwo hover:text-2xl cursor-pointer truncate'>
					{name}
				</a>
				<p className='text-sm mt-4 truncate'>{description}</p>
				<p className='text-sm mt-4 truncate'>{html_url}</p>
				<div className='flex gap-4'>
					<p className='text-sm'> Language : {language}</p>
					<p className='text-sm'> Stars - {stargazers_count}</p>
				</div>
			</div>
			<div className='max-w-fit'>
				<div className='mb-8 cursor-pointer'>
					<Badge color='bg-accentOne'>
						<a href={html_url} target='_blank'>
							{open_issues_count}
						</a>{" "}
					</Badge>
				</div>
				<div className='cursor-pointer'>
					<p className='text-accentTwo'>{pushed_at}</p>
				</div>
			</div>
		</div>
	);
};

export default Card;
