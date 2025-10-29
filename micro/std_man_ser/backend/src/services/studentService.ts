import prisma from "../prisma";

type ListOpts = {
  page?: number;
  limit?: number;
  search?: string;
  programId?: number;
  departmentId?: number;
  enrollmentYear?: number;
};

export async function createStudent(data: any) {
  // prevent duplicates by checking unique fields (email or studentNumber)
  const existing = await prisma.student.findFirst({
    where: {
      OR: [{ email: data.email }, { studentNumber: data.studentNumber }],
    },
  });
  if (existing) {
    const err: any = new Error(
      "Student with same email or student number already exists"
    );
    err.status = 409;
    err.code = "student.duplicate";
    throw err;
  }

  const created = await prisma.student.create({
    data: {
      studentNumber: data.studentNumber,
      firstName: data.firstName,
      lastName: data.lastName,
      email: data.email,
      dob: data.dob ? new Date(data.dob) : undefined,
      enrollmentYear: data.enrollmentYear,
      program: { connect: { id: Number(data.programId) } },
    },
    include: { program: true },
  });

  return created;
}

export async function listStudents(opts: ListOpts) {
  const page = opts.page || 1;
  const limit = Math.min(opts.limit || 20, 100);
  const where: any = {};

  if (opts.search) {
    where.OR = [
      { firstName: { contains: opts.search, mode: "insensitive" } },
      { lastName: { contains: opts.search, mode: "insensitive" } },
      { email: { contains: opts.search, mode: "insensitive" } },
      { studentNumber: { contains: opts.search, mode: "insensitive" } },
    ];
  }

  if (opts.programId) where.programId = opts.programId;
  if (opts.enrollmentYear) where.enrollmentYear = opts.enrollmentYear;

  // department filter: join through program -> department
  if (opts.departmentId) {
    where.program = { some: { departmentId: opts.departmentId } };
  }

  const [total, items] = await Promise.all([
    prisma.student.count({ where }),
    prisma.student.findMany({
      where,
      skip: (page - 1) * limit,
      take: limit,
      include: { program: true },
    }),
  ]);

  return { items, meta: { total, page, limit } };
}

export async function getStudentById(id: number) {
  return prisma.student.findUnique({
    where: { id },
    include: { program: true },
  });
}

export async function updateStudent(id: number, data: any) {
  // allow updating, but prevent unique conflicts
  if (data.email || data.studentNumber) {
    const conflict = await prisma.student.findFirst({
      where: {
        OR: [{ email: data.email }, { studentNumber: data.studentNumber }],
        NOT: { id },
      },
    });
    if (conflict) {
      const err: any = new Error("Conflict with existing student");
      err.status = 409;
      err.code = "student.duplicate";
      throw err;
    }
  }

  const updated = await prisma.student.update({
    where: { id },
    data: {
      firstName: data.firstName,
      lastName: data.lastName,
      email: data.email,
      dob: data.dob ? new Date(data.dob) : undefined,
      enrollmentYear: data.enrollmentYear,
      program: data.programId
        ? { connect: { id: Number(data.programId) } }
        : undefined,
    },
    include: { program: true },
  });

  return updated;
}

export async function deleteStudent(id: number) {
  await prisma.student.delete({ where: { id } });
}
