"use client";
import { useEffect, useState } from "react";
import StudentTable from "../../components/StudentTable";
import { fetchStudents } from "../../lib/api";

export default function StudentsPage() {
  const [students, setStudents] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState("");

  useEffect(() => {
    load();
  }, [page]);

  async function load() {
    try {
      const res = await fetchStudents({ page, limit: 20, search });
      setStudents(res.items || []);
    } catch (err) {
      // eslint-disable-next-line no-console
      console.error(err);
      setStudents([]);
    }
  }

  return (
    <div className="mx-auto max-w-4xl py-8">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-semibold">Students</h2>
        <a className="text-blue-600" href="/students/new">
          Add student
        </a>
      </div>

      <div className="mb-4 flex gap-2">
        <input
          className="border p-2 flex-1"
          placeholder="Search"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
        <button
          className="px-4 py-2 bg-gray-800 text-white"
          onClick={() => {
            setPage(1);
            load();
          }}
        >
          Search
        </button>
      </div>

      <StudentTable students={students} />

      <div className="mt-4 flex gap-2">
        <button
          className="px-3 py-1 border"
          onClick={() => setPage((p) => Math.max(1, p - 1))}
        >
          Prev
        </button>
        <span className="px-2">Page {page}</span>
        <button
          className="px-3 py-1 border"
          onClick={() => setPage((p) => p + 1)}
        >
          Next
        </button>
      </div>
    </div>
  );
}
