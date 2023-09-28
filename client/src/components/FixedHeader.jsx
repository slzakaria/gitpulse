import { FaGithub, FaSquareXTwitter } from "react-icons/fa6";

const FixedHeader = () => {
	return (
		<nav className='fixed top-0 inset-x-0 z-50 h-fit p-4 text-white bg-primary font-medium shadow-lg border-b-2 border-accentTwo'>
			<div className='p-2'>
				<div className='flex justify-between items-center'>
					<p className='text-textColor hover:text-accentTwo hover:underline'>GitPulse</p>
					<ul className='flex gap-4 mx-4 text-xl'>
						<li className='hover:scale-110'>
							<a
								href='https://github.com/Zackaria-Slimane'
								target='_blank'
								className='hover:text-accentTwo'>
								<FaGithub />
							</a>
						</li>
						<li className='hover:scale-110'>
							<a
								href='https://twitter.com/gitignorer'
								target='_blank'
								className='hover:text-accentTwo'>
								<FaSquareXTwitter />
							</a>
						</li>
					</ul>
				</div>
			</div>
		</nav>
	);
};

export default FixedHeader;
