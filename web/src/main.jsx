import { render } from 'preact'
import './index.css'
import Router from 'preact-router'
import { Header } from './core/components/header'
import { Footer } from './core/components/footer'
import { NotFound } from './pages/404'
import { Index } from './pages'
import { About } from './pages/about'
import { Article } from './pages/article'
import { DataProvider } from './core/providers/context'

const Main = () => {
  return (
    <DataProvider>
      <Header />
      <main className='flex-grow mt-12'>
        <Router>
          <Index path='/' element={<Index />} />
          <About path='/about' element={<About />} />
          <Article path='/articles/:category/:name' />
          <NotFound default />
        </Router>
      </main>
      <Footer />
    </DataProvider>
  )
}

render(<Main />, document.getElementById('root'))
