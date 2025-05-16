import { Link } from 'preact-router/match'

export const Header = () => {
  return (
    <header className='flex justify-between gap-1 md:mx-30 my-3 top-3 sticky'>
      <ul className='flex list-none m-0 p-0 justify-evenly md:gap-12 gap-3 items-center'>
        <li>
          <a href='https://github.com/loadept' target='_blank' title='github'
            className='hover:text-[#528bff] transition-colors'>
            <svg xmlns='http://www.w3.org/2000/svg' x='0px' y='0px' width='23' height='23' viewBox='0,0,256,256'
              className='hover:text-[#528bff] transition-colors'>
              <g fill='currentColor' fillRule='nonzero' stroke='none' strokeWidth='1' strokeLinecap='butt'
                strokeLinejoin='miter' strokeMiterlimit='10' strokeDasharray='' strokeDashoffset='0' fontFamily='none'
                fontWeight='none' fontSize='none' textAnchor='none' style={{ mixBlendMode: 'normal' }}>
                <g transform='scale(5.12,5.12)'>
                  <path
                    d='M17.791,46.836c0.711,-0.306 1.209,-1.013 1.209,-1.836v-5.4c0,-0.197 0.016,-0.402 0.041,-0.61c-0.014,0.004 -0.027,0.007 -0.041,0.01c0,0 -3,0 -3.6,0c-1.5,0 -2.8,-0.6 -3.4,-1.8c-0.7,-1.3 -1,-3.5 -2.8,-4.7c-0.3,-0.2 -0.1,-0.5 0.5,-0.5c0.6,0.1 1.9,0.9 2.7,2c0.9,1.1 1.8,2 3.4,2c2.487,0 3.82,-0.125 4.622,-0.555c0.934,-1.389 2.227,-2.445 3.578,-2.445v-0.025c-5.668,-0.182 -9.289,-2.066 -10.975,-4.975c-3.665,0.042 -6.856,0.405 -8.677,0.707c-0.058,-0.327 -0.108,-0.656 -0.151,-0.987c1.797,-0.296 4.843,-0.647 8.345,-0.714c-0.112,-0.276 -0.209,-0.559 -0.291,-0.849c-3.511,-0.178 -6.541,-0.039 -8.187,0.097c-0.02,-0.332 -0.047,-0.663 -0.051,-0.999c1.649,-0.135 4.597,-0.27 8.018,-0.111c-0.079,-0.5 -0.13,-1.011 -0.13,-1.543c0,-1.7 0.6,-3.5 1.7,-5c-0.5,-1.7 -1.2,-5.3 0.2,-6.6c2.7,0 4.6,1.3 5.5,2.1c1.699,-0.701 3.599,-1.101 5.699,-1.101c2.1,0 4,0.4 5.6,1.1c0.9,-0.8 2.8,-2.1 5.5,-2.1c1.5,1.4 0.7,5 0.2,6.6c1.1,1.5 1.7,3.2 1.6,5c0,0.484 -0.045,0.951 -0.11,1.409c3.499,-0.172 6.527,-0.034 8.204,0.102c-0.002,0.337 -0.033,0.666 -0.051,0.999c-1.671,-0.138 -4.775,-0.28 -8.359,-0.089c-0.089,0.336 -0.197,0.663 -0.325,0.98c3.546,0.046 6.665,0.389 8.548,0.689c-0.043,0.332 -0.093,0.661 -0.151,0.987c-1.912,-0.306 -5.171,-0.664 -8.879,-0.682c-1.665,2.878 -5.22,4.755 -10.777,4.974v0.031c2.6,0 5,3.9 5,6.6v5.4c0,0.823 0.498,1.53 1.209,1.836c9.161,-3.032 15.791,-11.672 15.791,-21.836c0,-12.682 -10.317,-23 -23,-23c-12.683,0 -23,10.318 -23,23c0,10.164 6.63,18.804 15.791,21.836z'>
                  </path>
                </g>
              </g>
            </svg>
          </a>
        </li>
        <li>
          <a href='https://www.linkedin.com/in/jesus-machaca-hancco-136119276' target='_blank' title='linkedin'
            className='hover:text-[#528bff] transition-colors'>
            <svg xmlns='http://www.w3.org/2000/svg' x='0px' y='0px' width='23' height='23' viewBox='0,0,256,256'
              className='hover:text-[#528bff] transition-colors'>
              <g fill='currentColor' fillRule='nonzero' stroke='none' strokeWidth='1' strokeLinecap='butt'
                strokeLinejoin='miter' strokeMiterlimit='10' strokeDasharray='' strokeDashoffset='0' fontFamily='none'
                fontWeight='none' fontSize='none' textAnchor='none' style={{ mixBlendMode: 'normal' }}>
                <g transform='scale(5.12,5.12)'>
                  <path
                    d='M41,4h-32c-2.76,0 -5,2.24 -5,5v32c0,2.76 2.24,5 5,5h32c2.76,0 5,-2.24 5,-5v-32c0,-2.76 -2.24,-5 -5,-5zM17,20v19h-6v-19zM11,14.47c0,-1.4 1.2,-2.47 3,-2.47c1.8,0 2.93,1.07 3,2.47c0,1.4 -1.12,2.53 -3,2.53c-1.8,0 -3,-1.13 -3,-2.53zM39,39h-6c0,0 0,-9.26 0,-10c0,-2 -1,-4 -3.5,-4.04h-0.08c-2.42,0 -3.42,2.06 -3.42,4.04c0,0.91 0,10 0,10h-6v-19h6v2.56c0,0 1.93,-2.56 5.81,-2.56c3.97,0 7.19,2.73 7.19,8.26z'>
                  </path>
                </g>
              </g>
            </svg>
          </a>
        </li>
        <li>
          <a href='https://www.tiktok.com/@loadept' target='_blank' title='tiktok'
            className='hover:text-[#528bff] transition-colors'>
            <svg xmlns='http://www.w3.org/2000/svg' x='0px' y='0px' width='23' height='23' viewBox='0,0,256,256'
              className='hover:text-[#528bff] transition-colors'>
              <g fill='currentColor' fillRule='nonzero' stroke='none' strokeWidth='1' strokeLinecap='butt'
                strokeLinejoin='miter' strokeMiterlimit='10' strokeDasharray='' strokeDashoffset='0' fontFamily='none'
                fontWeight='none' fontSize='none' textAnchor='none' style={{ mixBlendMode: 'normal' }}>
                <g transform='scale(5.12,5.12)'>
                  <path
                    d='M41,4h-32c-2.757,0 -5,2.243 -5,5v32c0,2.757 2.243,5 5,5h32c2.757,0 5,-2.243 5,-5v-32c0,-2.757 -2.243,-5 -5,-5zM37.006,22.323c-0.227,0.021 -0.457,0.035 -0.69,0.035c-2.623,0 -4.928,-1.349 -6.269,-3.388c0,5.349 0,11.435 0,11.537c0,4.709 -3.818,8.527 -8.527,8.527c-4.709,0 -8.527,-3.818 -8.527,-8.527c0,-4.709 3.818,-8.527 8.527,-8.527c0.178,0 0.352,0.016 0.527,0.027v4.202c-0.175,-0.021 -0.347,-0.053 -0.527,-0.053c-2.404,0 -4.352,1.948 -4.352,4.352c0,2.404 1.948,4.352 4.352,4.352c2.404,0 4.527,-1.894 4.527,-4.298c0,-0.095 0.042,-19.594 0.042,-19.594h4.016c0.378,3.591 3.277,6.425 6.901,6.685z'>
                  </path>
                </g>
              </g>
            </svg>
          </a>
        </li>
        <li>
          <a href='https://www.youtube.com/@load3pt' target='_blank' title='youtube'>
            <svg xmlns='http://www.w3.org/2000/svg' x='0px' y='0px' width='23' height='23' viewBox='0,0,256,256'
              className='hover:text-[#528bff] transition-colors'>
              <g fill='currentColor' fillRule='nonzero' stroke='none' strokeWidth='1' strokeLinecap='butt'
                strokeLinejoin='miter' strokeMiterlimit='10' strokeDasharray='' strokeDashoffset='0' fontFamily='none'
                fontWeight='none' fontSize='none' textAnchor='none' style={{ mixBlendMode: 'normal' }}>
                <g transform='scale(8.53333,8.53333)'>
                  <path
                    d='M15,4c-4.186,0 -9.61914,1.04883 -9.61914,1.04883l-0.01367,0.01563c-1.90652,0.30491 -3.36719,1.94317 -3.36719,3.93555v6v0.00195v5.99805v0.00195c0.00384,1.96564 1.4353,3.63719 3.37695,3.94336l0.00391,0.00586c0,0 5.43314,1.05078 9.61914,1.05078c4.186,0 9.61914,-1.05078 9.61914,-1.05078l0.00195,-0.00195c1.94389,-0.30554 3.37683,-1.97951 3.37891,-3.94727v-0.00195v-5.99805v-0.00195v-6c-0.00288,-1.96638 -1.43457,-3.63903 -3.37695,-3.94531l-0.00391,-0.00586c0,0 -5.43314,-1.04883 -9.61914,-1.04883zM12,10.39844l8,4.60156l-8,4.60156z'>
                  </path>
                </g>
              </g>
            </svg>
          </a>
        </li>
      </ul>
      <ul className='flex text-xs md:text-sm list-none m-0 p-0 justify-evenly md:gap-12 gap-3 items-center'>
        <li className='hover:text-[#528bff] transition-colors'>
          <Link href='/' title='home'>
            Inicio
          </Link>
        </li>
        <li className='hover:text-[#528bff] transition-colors'>
          <Link href='/about' title='about me'>
            Sobre mí
          </Link>
        </li>
        <li className='hover:text-[#528bff] transition-colors'>
          Recursos
        </li>
        <li className='hover:text-[#528bff] transition-colors'>
          Contacto
        </li>
      </ul>
    </header>
  )
}
