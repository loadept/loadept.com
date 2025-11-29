import type { Category, Post, Resources } from '../types/home'
import { ResourcesTree } from './ResourcesTree'

interface Props {
  categories: Category[]
  posts: Post[]
  resources: Resources[]
}

export const Content = ({ categories, posts, resources }: Props) => {
  return (
    <section className="space-y-6 mt-15 mb-30">
      <div className="flex items-center gap-3">
        <span className="text-3xl text-[#98c379]"></span>
        <h2 className="text-2xl font-bold text-[#e5c07b]">Contenido</h2>
      </div>
      <ResourcesTree resources={resources} categories={categories} posts={posts} />
    </section>
  )
}
