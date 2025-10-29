import { Request, Response, NextFunction } from 'express';
import * as studentService from '../services/studentService';

export async function createStudent(req: Request, res: Response, next: NextFunction) {
  try {
    const payload = req.body;
    const student = await studentService.createStudent(payload);
    res.status(201).json({ ok: true, data: student });
  } catch (err) {
    next(err);
  }
}

export async function listStudents(req: Request, res: Response, next: NextFunction) {
  try {
    const { page, limit, search, programId, departmentId, enrollmentYear } = req.query;
    const result = await studentService.listStudents({
      page: Number(page) || 1,
      limit: Number(limit) || 20,
      search: typeof search === 'string' ? search : undefined,
      programId: programId ? Number(programId) : undefined,
      departmentId: departmentId ? Number(departmentId) : undefined,
      enrollmentYear: enrollmentYear ? Number(enrollmentYear) : undefined
    });
    res.json({ ok: true, data: result.items, meta: result.meta });
  } catch (err) {
    next(err);
  }
}

export async function getStudent(req: Request, res: Response, next: NextFunction) {
  try {
    const id = Number(req.params.id);
    const student = await studentService.getStudentById(id);
    if (!student) return res.status(404).json({ ok: false, error: { code: 'student.not_found', message: 'Student not found' } });
    res.json({ ok: true, data: student });
  } catch (err) {
    next(err);
  }
}

export async function updateStudent(req: Request, res: Response, next: NextFunction) {
  try {
    const id = Number(req.params.id);
    const payload = req.body;
    const student = await studentService.updateStudent(id, payload);
    res.json({ ok: true, data: student });
  } catch (err) {
    next(err);
  }
}

export async function deleteStudent(req: Request, res: Response, next: NextFunction) {
  try {
    const id = Number(req.params.id);
    await studentService.deleteStudent(id);
    res.status(204).send();
  } catch (err) {
    next(err);
  }
}
