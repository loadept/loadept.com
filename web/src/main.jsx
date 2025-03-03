import { render } from 'preact'
import Router from 'preact-router'
import './index.css'
import { Header } from './core/components/header'
import { Footer } from './core/components/footer'
import { Index } from './pages/index'
import { About } from './pages/about'
import { NotFound } from './pages/404'

const Main = () => {
  return (
    <>
      <Header />
      <main class="flex-grow mt-12">
        <Router>
          <Index path='/' />
          <About path='/about' />
          <NotFound default />
        </Router>
      </main>
      <Footer />
    </>
  )
}

render(<Main />, document.getElementById('app'))
