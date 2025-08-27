export const DEFAULT_BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

export const ROUTES = {
  COMPONENTS: `${DEFAULT_BACKEND_URL}/components`,
  COMPONENT_CATEGORY: (category: string, page: string) => `${DEFAULT_BACKEND_URL}/components/${category}?page=${page}`,
  COMPONENT_BRAND: (category: string, brand: string, page: string) =>
    `${DEFAULT_BACKEND_URL}/components/${category}/${brand}?page=${page}`,
  COMPONENT_ID: (id: string, page: string) =>
    `${DEFAULT_BACKEND_URL}/components/item/${id}?page=${page}`,
};
