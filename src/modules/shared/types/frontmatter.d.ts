import type { Category } from '../../home/types/home'

export interface FrontMatter {
  title: string
  date: string
  tags: string[]
  keywords: string[]
  category: Category
}

export interface FrontMatterResource {
  layout: string
  title: string
  description: string
  keywords: string[]
  goImport: string
  goSource: string
}
