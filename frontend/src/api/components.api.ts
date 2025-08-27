import { ROUTES } from "@/constants/routes.constants";
import { Component, getComponentByIdRequest, getComponentsByBrandRequest, getComponentsByCategoryRequest, getComponentsRequest } from "@/types/components.type";

export const getComponents = async ({ page = "1" }: getComponentsRequest): Promise<Component[]> => {
  const response = await fetch(ROUTES.COMPONENTS + `?page=${page}`);
  const data = await response.json();
  
  // Handle case where data is not an array (e.g., error response)
  if (!Array.isArray(data)) {
    return [data];
  }
  
  return data.map((component: Component) => ({
    ...component,
    id: component.id.toString(),
  }));
};

export const getComponentsByCategory = async ({page = "1", category}: getComponentsByCategoryRequest): Promise<Component[]> => {
  const response = await fetch(ROUTES.COMPONENT_CATEGORY(category, page));
  const data = await response.json();
  
  // Handle case where data is not an array (e.g., error response)
  if (!Array.isArray(data)) {
    return [data];
  }
  
  return data.map((component: Component) => ({
    ...component,
    id: component.id.toString(),
  }));
};

export const getComponentsByBrand = async ({page = "1", category, brand}: getComponentsByBrandRequest): Promise<Component[]> => {
  const response = await fetch(ROUTES.COMPONENT_BRAND(category, brand, page));
  const data = await response.json();
  
  // Handle case where data is not an array (e.g., error response)
  if (!Array.isArray(data)) {
    return [data];
  }
  
  return data.map((component: Component) => ({
    ...component,
    id: component.id.toString(),
  }));
};

export const getComponentById = async ({page = "1", id}: getComponentByIdRequest): Promise<Component> => {
  const response = await fetch(ROUTES.COMPONENT_ID(id, page));
  const data = await response.json();
  return {
    ...data,
    id: data.id?.toString() || id,
  };
};
