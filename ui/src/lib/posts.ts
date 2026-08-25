import type { CollectionEntry } from 'astro:content'

export type Post = CollectionEntry<'posts'>['data'] & {
  slug: string
  category: string
}

export const toPost = ({ id, data }: CollectionEntry<'posts'>): Post => ({
  slug: id,
  category: id.split('/')[0].replaceAll('-', ' '),
  ...data,
})
