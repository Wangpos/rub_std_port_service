"use client";
import Link from "next/link";

export default function StudentTable({ students }: { students: any[] }) {
  return (
    <table className="min-w-full table-auto border-collapse">
      <thead>
        <tr>
          <th className="text-left border-b p-2">#</th>
          <th className="text-left border-b p-2">Name</th>
          <th className="text-left border-b p-2">Email</th>
          <th className="text-left border-b p-2">Program</th>
          <th className="text-left border-b p-2">Actions</th>
        </tr>
      </thead>
      <tbody>
        {students.map((s) => (
          <tr key={s.id} className="hover:bg-gray-50">
            <td className="p-2">{s.studentNumber}</td>
            <td className="p-2">
              {s.firstName} {s.lastName}
            </td>
            <td className="p-2">{s.email}</td>
            <td className="p-2">{s.program?.name ?? "-"}</td>
            <td className="p-2">
              <Link className="text-blue-600" href={`/students/${s.id}`}>
                Edit
              </Link>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
