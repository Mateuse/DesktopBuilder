export type Component = {
  id: string;
  category: string;
  brand: string;
  model: string;
  sku: string;
  upc: string;
  specs: string;
  releaseDate: string;
  createdAt: string;
};

export type Page = {
  page?: string;
}

export type getComponentsRequest = Page;

export type getComponentsByCategoryRequest = Page & {
  category: string;
};

export type getComponentsByBrandRequest = Page & {
  category: string;
  brand: string;
};

export type getComponentByIdRequest = Page & {
  id: string;
};
