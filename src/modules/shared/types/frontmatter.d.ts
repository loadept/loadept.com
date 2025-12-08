export interface FrontMatter {
  title: string
  date: string
  keywords?: string[]
  category: string
}

export interface FrontMatterResource {
  layout: string
  title: string
  description: string
  keywords: string[]
  goImport: string
  goSource: string
}
