export interface Post {
  slug: string
  title: string
  date: string
  category: Category
}

export interface Category {
  name: string
  color: string
  icon: string
}

export interface ResourceItem {
  label: string
  href?: string
}

export interface Resources {
  label: string
  items: ResourceItem[]
}
