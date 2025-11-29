import { ChevronDown, ChevronRight } from 'lucide-preact'
import { useState } from 'preact/hooks'
import type { Category, Post, Resources } from '../types/home'

export const ResourcesTree = (
  { resources, categories, posts }:
  { resources: Resources[], categories: Category[], posts: Post[] }
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
            className="flex items-center gap-2 text-[#abb2bf] hover:text-[#61afef] transition-colors group w-full"
          >
            {expandedResources.includes(resource.label) ? (
              <ChevronDown className="h-5 w-5 text-[#98c379]" />
            ) : (
              <ChevronRight className="h-5 w-5 text-[#98c379]" />
            )}
            <span className="text-[#98c379] text-lg">|</span>
            <span className="font-mono text-lg font-semibold">
              {resource.label}
            </span>
          </button>

          {expandedResources.includes(resource.label) && (
            <div className="ml-8 mt-3 space-y-2 border-l border-[#98c379] pl-4">
              {resource.label === "Posts"
                ? categories.map((category, k) => {
                    const filteredPosts = posts.filter((a) => a.category.name === category.name)

                    return (
                      <div key={k}>
                        <button
                          onClick={() => toggleArticleCategory(category.name)}
                          className="flex items-center gap-2 text-[#abb2bf] hover:text-[#61afef] transition-colors group"
                        >
                          {expandedArticleCategories.includes(category.name) ? (
                            <ChevronDown className="h-4 w-4 text-[#61afef]" />
                          ) : (
                            <ChevronRight className="h-4 w-4 text-[#61afef]" />
                          )}
                          <span className="ml-1 text-[#3e4451] text-base">├──</span>
                          <span className="font-mono text-base">{category.name}</span>
                        </button>

                        {expandedArticleCategories.includes(category.name) && (
                          <div className="ml-8 mt-2 space-y-1 border-l border-[#3e4451] pl-4">
                            {filteredPosts.map((post, k) => (
                              <a
                                key={k}
                                href={`/posts/${post.slug}`}
                                className="flex items-center gap-2 text-[#abb2bf] hover:text-[#61afef] transition-colors"
                              >
                                <span className="text-[#3e4451] text-sm">└──</span>
                                <span
                                  className="font-mono text-sm"
                                  style={{
                                    viewTransitionName: `post-${post.slug.replace(/\//g, '-').replace(/ /g, '-')}`,
                                  }}
                                >
                                  {post.title}
                                </span>
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
                      className={`flex items-center gap-2 text-[#abb2bf] transition-colors group ${
                        item.href === "#" ? "cursor-not-allowed opacity-50" : "hover:text-[#61afef]"
                      }`}
                    >
                      <span className="text-[#3e4451] text-lg">├──</span>
                      {item.href === "#" && <span className="text-[#3e4451] text-sm">//</span>}
                      <span className="font-mono">{item.label}</span>
                    </a>
                  ))}
            </div>
          )}
        </div>
      ))}
    </div>
  )
}
