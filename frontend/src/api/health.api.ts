// Checks health of backend services
import { ROUTES } from "../constants/urls"
import { HealthResponseSchema, type HealthResponse } from "../schemas/health.schemas"
import { createValidatedFetch } from "../utils/validation.utils"

// Golang main backend
export const mainHealth = async (): Promise<HealthResponse> => {
    return createValidatedFetch(HealthResponseSchema)(ROUTES.MAIN_BE_HEALTH);
}
