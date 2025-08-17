import { Container, Title, Text } from "@mantine/core";
import { mainHealth } from "../api/health.api";
import { useQuery } from "@tanstack/react-query";

export const HealthComponent = () => {

    const { data, isLoading, error } = useQuery({
        queryKey: ["health"],
        queryFn: mainHealth,
    });

    if (isLoading) return <Text>Loading...</Text>

    if (error) return <Text>Error: {error instanceof Error ? error.message : 'Unknown error'}</Text>

    return (
        <Container>
            <Title order={1}>Health</Title>
            <Text>{data?.message || 'No message available'}</Text>
        </Container>
    )
}