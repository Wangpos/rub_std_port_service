const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:4000";

export async function fetchStudents(
  params: {
    page?: number;
    limit?: number;
    search?: string;
    programId?: number;
    enrollmentYear?: number;
  } = {}
) {
  const qs = new URLSearchParams();
  if (params.page) qs.set("page", String(params.page));
  if (params.limit) qs.set("limit", String(params.limit));
  if (params.search) qs.set("search", params.search);
  if (params.programId) qs.set("programId", String(params.programId));
  if (params.enrollmentYear)
    qs.set("enrollmentYear", String(params.enrollmentYear));

  const res = await fetch(`${API_URL}/api/students?${qs.toString()}`);
  if (!res.ok) throw new Error("Failed to fetch students");
  return res.json();
}

export async function createStudent(payload: any) {
  const res = await fetch(`${API_URL}/api/students`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error("Create failed");
  return res.json();
}

export async function getStudent(id: number) {
  const res = await fetch(`${API_URL}/api/students/${id}`);
  if (!res.ok) throw new Error("Failed to fetch student");
  return res.json();
}

export async function updateStudent(id: number, payload: any) {
  const res = await fetch(`${API_URL}/api/students/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error("Update failed");
  return res.json();
}

export async function deleteStudent(id: number) {
  const res = await fetch(`${API_URL}/api/students/${id}`, {
    method: "DELETE",
  });
  if (res.status !== 204) throw new Error("Delete failed");
  return true;
}

export async function fetchPrograms() {
  const res = await fetch(`${API_URL}/api/colleges`);
  if (!res.ok) throw new Error("Failed to fetch entities");
  return res.json();
}
