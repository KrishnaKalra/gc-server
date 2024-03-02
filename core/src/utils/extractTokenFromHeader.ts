import { Request } from 'express';

export function extractTokenFromHeader(req:any): string | null {
  // Get the Authorization header from the request
  const authHeader = req.headers.authorization;

  // Check if Authorization header is present and in the format 'Bearer <token>'
  if (authHeader && authHeader.startsWith('Bearer ')) {
    // Extract the token part after 'Bearer '
    const token = authHeader.substring('Bearer '.length);
    return token;
  } else {
    // If Authorization header is not present or not in the correct format, return null
    return null;
  }
}
