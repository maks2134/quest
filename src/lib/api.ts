import type { Procedure } from "@/types/procedure";

const API_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:8000/api/v1";

export async function fetchProcedures(): Promise<Procedure[]> {
  const response = await fetch(`${API_BASE_URL}/procedures`);
  if (!response.ok) {
    throw new Error(`Failed to fetch procedures: ${response.statusText}`);
  }
  
  return response.json();
}

