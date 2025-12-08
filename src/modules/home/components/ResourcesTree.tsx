import { useState } from 'preact/hooks'
import type { Post, Resources } from '../types/home'

export const ResourcesTree = (
  { resources, categories, posts }:
  { resources: Resources[], categories: string[], posts: Post[] }
) => {
  const [expandedResources, setExpandedResources] = useState<string[]>(['Posts', 'Paquetes Go', 'Más herramientas'])
  const [expandedArticleCategories, setExpandedArticleCategories] = useState<string[]>([])

  const toggleResource = (resource: string) => {
    setExpandedResources((prev) =>
      prev.includes(resource) ? prev.filter((c) => c !== resource) : [...prev, resource],
    )
  }

  const toggleArticleCategory = (categoryId: string) => {
    setExpandedArticleCategories((prev) =>
      prev.includes(categoryId) ? prev.filter((c) => c !== categoryId) : [...prev, categoryId],
    )
  }

  return (
    <div className="space-y-4 text-base">
      {resources.map((resource) => (
        <div key={resource.label}>
          <button
            onClick={() => toggleResource(resource.label)}
            className="flex items-center gap-2 text-[#61afef] hover:text-[#528bff] transition-colors w-full font-bold text-lg cursor-pointer"
          >
            <span>{expandedResources.includes(resource.label) ? "▼" : "▶"}</span>
            <span>|</span>
            <span>{resource.label}</span>
          </button>

          {expandedResources.includes(resource.label) && (
            <div className="ml-6 mt-3 space-y-2 border-l border-[#61afef] pl-4">
              {resource.label === "Posts"
                ? categories.map((category, k) => {
                    const filteredPosts = posts.filter((a) => a.category === category)

                    return (
                      <div key={k}>
                        <button
                          onClick={() => toggleArticleCategory(category)}
                          className="flex items-center gap-2 text-[#abb2bf] hover:text-[#c5c8c6] transition-colors font-semibold text-base cursor-pointer"
                        >
                          <span>{expandedArticleCategories.includes(category) ? "▼" : "▶"}</span>
                          <span>├──</span>
                          <span>{category}</span>
                        </button>

                        {expandedArticleCategories.includes(category) && (
                          <div className="ml-6 mt-2 space-y-1 border-l border-[#abb2bf] pl-4">
                            {filteredPosts.map((post, k) => (
                              <a
                                key={k}
                                href={`/posts/${post.slug}`}
                                className="flex items-center gap-2 text-[#abb2bf] hover:text-[#528bff] transition-colors text-sm hover:underline"
                              >
                                <span>└──</span>
                                <span>{post.title}</span>
                              </a>
                            ))}
                          </div>
                        )}
                      </div>
                    )
                  })
                : resource.items.map((item) => (
                    <a
                      key={item.label}
                      href={item.href || "#"}
                      className={`flex items-center gap-2 text-base transition-colors ${
                        item.href === "#"
                          ? "text-[#5c6370] cursor-not-allowed opacity-60"
                          : "text-[#abb2bf] hover:text-[#528bff]"
                      }`}
                    >
                      <span>├──</span>
                      <span>{item.label}</span>
                    </a>
                  ))}
            </div>
          )}
        </div>
      ))}
    </div>
  )
}
