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
        'Content-Type': 'application/json'  // Menambahkan Content-Type JSON
      }
    });

    // Debug response jika terjadi kesalahan
    console.log("Response status:", response.status());
    console.log("Response text:", await response.text());

    // Periksa status harus 201 jika berhasil
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
        'Content-Type': 'application/json'  // Menambahkan Content-Type JSON
      }
    });

    // Debug response jika terjadi kesalahan
    console.log("Response status:", response.status());
    console.log("Response text:", await response.text());

    // Periksa status harus 400 jika gagal
    expect(response.status()).toBe(400);
  });
});
