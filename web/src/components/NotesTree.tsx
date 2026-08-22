import { useState } from 'preact/hooks'
import type { Post } from '../lib/posts'

interface Props {
  posts: Post[]
}

export const NotesTree = ({ posts }: Props) => {
  const categories = [...new Set(posts.map((post) => post.category))]
  const [expandedCategories, setExpandedCategories] = useState<string[]>([])

  const toggleCategory = (categoryId: string) => {
    setExpandedCategories((prev) =>
      prev.includes(categoryId) ? prev.filter((c) => c !== categoryId) : [...prev, categoryId],
    )
  }

  return (
    <div>
      <h2 className="text-primary font-normal text-lg">Notas</h2>
      <ul className="ml-6 mt-2 space-y-1">
        {categories.map((category) => {
          const categoryPosts = posts.filter((post) => post.category === category)
          return (
            <li key={category}>
              <button
                onClick={() => toggleCategory(category)}
                className="text-foreground hover:text-primary transition-colors font-normal cursor-pointer"
              >
                {expandedCategories.includes(category) ? "▼" : "▶"} {category}
              </button>

              {expandedCategories.includes(category) && (
                <ul className="ml-6 mt-2 space-y-1">
                  {categoryPosts.map((post) => (
                    <li key={post.slug}>
                      <a href={`/posts/${post.slug}`} className="hover:underline">
                        {post.title}
                      </a>
                    </li>
                  ))}
                </ul>
              )}
            </li>
          )
        })}
      </ul>
    </div>
  )
}
