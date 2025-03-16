import { createRoot } from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'
import { Header } from './core/components/header'
import { Footer } from './core/components/footer'
import { NotFound } from './pages/404'
import { Index } from './pages'
import { About } from './pages/about'
import { Article } from './pages/article'
import './index.css'

createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Header />
    <main className='flex-grow mt-12'>
      <Routes>
        <Route path='/' element={<Index />} />
        <Route path='/about' element={<About />} />
        <Route path='/article/:category/:name' element={<Article />} />
        <Route path='*' element={<NotFound />} />
      </Routes>
    </main>
    <Footer />
  </BrowserRouter>
)
