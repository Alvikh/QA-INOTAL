const { test, expect } = require('@playwright/test');

// URL dasar API
const BASE_URL = 'http://127.0.0.1:8080/sistem/v010/aplikasi';

// Credentials for Basic Auth
const USERNAME = 'user-rsud';
const PASSWORD = 'password123';

test.describe('Aplikasi API Tests', () => {
  // Test: Create Aplikasi dengan data valid
  test('Create Aplikasi with valid data', async ({ request }) => {
    const response = await request.post(`${BASE_URL}`, {
      data: {
        "nama": "Aplikasi Test",
        "label": "Test Label",
        "logo": "logo.png",
        "url_fe": "http://example.com",
        "url_api": "http://api.example.com"
      },
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`,
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(201);
  });

  // Test: Create Aplikasi dengan data invalid
  test('Create Aplikasi with invalid data', async ({ request }) => {
    const response = await request.post(`${BASE_URL}`, {
      data: {
        "nama": "", // nama kosong
        "label": "Test Label",
        "logo": "logo.png",
        "url_fe": "http://example.com",
        "url_api": "http://api.example.com"
      },
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`,
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(400);
  });

  // Test: Update Aplikasi dengan kd valid dan data valid
  test('Update Aplikasi with valid kd and data', async ({ request }) => {
    const kd = 1; // kd valid
    const response = await request.put(`${BASE_URL}/${kd}`, {
      data: {
        "nama": "Aplikasi Test Update",
        "label": "Test Label Update",
        "logo": "logoUpdate.png",
        "url_fe": "http://example.com",
        "url_api": "http://api.example.com"
      },
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`,
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: Update Aplikasi dengan data tidak valid
  test('Update Aplikasi with invalid data', async ({ request }) => {
    const kd = 1;
    const response = await request.put(`${BASE_URL}/${kd}`, {
      data: {
        "nama": "", // nama kosong
        "label": "Invalid Label",
        "logo": "invalid_logo.png",
        "url_fe": "http://invalid.com",
        "url_api": "http://api.invalid.com"
      },
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`,
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(400);
  });

  // Test: Delete Aplikasi dengan kd valid
  test('Delete Aplikasi with valid kd', async ({ request }) => {
    const kd = 1;
    const response = await request.delete(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: Delete Aplikasi dengan kd tidak valid
  test('Delete Aplikasi with invalid kd', async ({ request }) => {
    const kd = 9999; // kd tidak valid
    const response = await request.delete(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(404);
  });

  // Test: FindAll pada Aplikasi
  test('FindAll Aplikasi', async ({ request }) => {
    const response = await request.get(`${BASE_URL}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: FindByLimit (limit=5) pada Aplikasi
  test('FindByLimit Aplikasi with limit=5', async ({ request }) => {
    const limit = 2;
    const response = await request.get(`${BASE_URL}?limit=${limit}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: FindByKd pada Aplikasi dengan kd valid
  test('FindByKd Aplikasi with valid kd', async ({ request }) => {
    const kd = 2; // kd valid
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: FindByKd pada Aplikasi dengan kd tidak valid
  test('FindByKd Aplikasi with invalid kd', async ({ request }) => {
    const kd = 9999; // kd tidak valid
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(404);
  });

  test('Create Aplikasi without authentication', async ({ request }) => {
    const response = await request.post(`${BASE_URL}`, {
      data: {
        "nama": "Aplikasi Tanpa Autentikasi",
        "label": "Label Tanpa Auth",
        "logo": "logo.png",
        "url_fe": "http://example.com",
        "url_api": "http://api.example.com"
      },
      headers: {
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(401); // Unauthorized
  });

  test('SQL Injection attempt in nama field', async ({ request }) => {
    const sqlInjectionPayload = "'; DROP TABLE Aplikasi; --";
    
    const response = await request.post(`${BASE_URL}`, {
      data: {
        "nama": sqlInjectionPayload,
        "label": "Test Label",
        "logo": "logo.png",
        "url_fe": "http://example.com",
        "url_api": "http://api.example.com"
      },
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`,
        'Content-Type': 'application/json'
      }
    });
  
    // Expect a 400 Bad Request status
    expect(response.status()).toBe(400);
  
    // Log and check the structure of the response
    const responseBody = await response.json();
    console.log(responseBody); // Debug log for inspecting the response body
  
    // Update the check to match the actual error message
    expect(responseBody.error).toContain("Invalid input data (SQL Injection detected)");
  });
});
