const SideBar = () => {
	return (
		<aside className='hidden pt-16 w-[35dvw] border-r-2 border-accentTwo text-white bg-primary fixed inset-y-0 overflow-x-hidden overflow-hidden sm:block'>
			<div className='p-4 min-h-full mt-4'>
				<div className='px-6 pb-6 flex items-start sm:flex-col'>
					<div className='hidden px-3 text-white text-base sm:block'>
						<h2 className='text-white underline my-2'>About GitPulse :</h2>
						<p className='text-gray-400'>
							GitPulse fetches and lists all github repositories who have open issues and have been
							active the last 3 months.
							<br />
							Find active easy repos to help on !
						</p>
					</div>
					<div className='hidden px-3 text-white text-base sm:block mt-10'>
						<h2 className='text-white underline my-2'>Filter by technology :</h2>
						<p className='text-gray-400'>work in progress</p>
					</div>
				</div>
			</div>
		</aside>
	);
};

export default SideBar;
