const { PrismaClient } = require("@prisma/client");
const prisma = new PrismaClient();

async function main() {
  console.log("Seeding database...");

  // Create a college
  const college = await prisma.college.create({
    data: { name: "Example College" },
  });
  console.log("Created college", college.id);

  // Create a department
  const department = await prisma.department.create({
    data: { name: "Computer Science", collegeId: college.id },
  });
  console.log("Created department", department.id);

  // Create a program
  const program = await prisma.program.create({
    data: {
      name: "BSc Computer Science",
      departmentId: department.id,
      collegeId: college.id,
    },
  });
  console.log("Created program", program.id);

  // Create some students
  const students = [
    {
      studentNumber: "S2025001",
      firstName: "Alice",
      lastName: "Anderson",
      email: "alice@example.edu",
      enrollmentYear: 2025,
      programId: program.id,
    },
    {
      studentNumber: "S2025002",
      firstName: "Bob",
      lastName: "Brown",
      email: "bob@example.edu",
      enrollmentYear: 2025,
      programId: program.id,
    },
  ];

  for (const s of students) {
    const created = await prisma.student.create({ data: s });
    console.log("Created student", created.id);
  }

  console.log("Seeding finished.");
}

main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    await prisma.$disconnect();
  });
