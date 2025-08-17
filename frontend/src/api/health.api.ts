// Checks health of backend services
import { ROUTES } from "../constants/urls"
import type { HealthResponse } from "../types/api.types"

// Golang main backend
export const mainHealth = async (): Promise<HealthResponse> => {
    const response = await fetch(ROUTES.MAIN_BE_HEALTH)
    if (!response.ok) {
        throw new Error("Failed to fetch health")
    }
    const data = await response.json() as HealthResponse;
    return data;
}