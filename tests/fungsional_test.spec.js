const { test, expect } = require('@playwright/test');

// URL dasar API
const BASE_URL = 'http://127.0.0.1:8080/sistem/v010/aplikasi';

// Credentials for Basic Auth
const USERNAME = 'user-rsud';
const PASSWORD = 'password123';

test.describe('Aplikasi API Tests - Additional Scenarios', () => {

  // Test: Validasi endpoint untuk mengambil data aplikasi berdasarkan kd dengan data valid
  test('Find Aplikasi by valid kd', async ({ request }) => {
    const kd = 2; // kd valid
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: Validasi endpoint untuk mengambil data aplikasi berdasarkan kd yang tidak ada di database
  test('Find Aplikasi by invalid kd', async ({ request }) => {
    const kd = 9999; // kd tidak ada
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(404);
  });

  // Test: Validasi endpoint untuk mengambil data aplikasi dengan format kd yang tidak valid (alfanumerik)
  test('Find Aplikasi by non-numeric kd', async ({ request }) => {
    const kd = 'abc123'; // kd alfanumerik
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(400); // Bad Request karena format kd salah
  });

  // Test: Validasi endpoint untuk mengambil data aplikasi tanpa token otentikasi
  test('Find Aplikasi without authentication', async ({ request }) => {
    const kd = 1; // kd valid
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(401); // Unauthorized
  });

  // Test: Validasi respon waktu (response time) untuk endpoint mengambil data aplikasi
  test('Response time for Find Aplikasi by valid kd', async ({ request }) => {
    const kd = 2; // kd valid
    const response = await request.get(`${BASE_URL}/${kd}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
    expect(response.time()).toBeLessThan(1000); // Respon dalam waktu kurang dari 1000ms
  });

  // Test: Validasi FindByLimit dengan limit=5
  test('FindByLimit with limit=5', async ({ request }) => {
    const limit = 5;
    const response = await request.get(`${BASE_URL}?limit=${limit}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: Validasi FindByLimit dengan limit melebihi jumlah data tersedia
  test('FindByLimit with limit exceeding available data', async ({ request }) => {
    const limit = 1000; // limit lebih besar dari data yang tersedia
    const response = await request.get(`${BASE_URL}?limit=${limit}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200); // Data yang ada tetap dikembalikan
    expect(response.body().length).toBeLessThanOrEqual(limit); // Jumlah data yang dikembalikan tidak melebihi limit
  });

  // Test: Validasi FindAll pada Aplikasi
  test('FindAll Aplikasi', async ({ request }) => {
    const response = await request.get(`${BASE_URL}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
  });

  // Test: Validasi FindAll tanpa autentikasi
  test('FindAll Aplikasi without authentication', async ({ request }) => {
    const response = await request.get(`${BASE_URL}`, {
      headers: {
        'Content-Type': 'application/json'
      }
    });
    expect(response.status()).toBe(401); // Unauthorized
  });

  // Test: Validasi respon waktu (response time) untuk FindAll endpoint
  test('Response time for FindAll Aplikasi', async ({ request }) => {
    const response = await request.get(`${BASE_URL}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
    expect(response.time()).toBeLessThan(1000); // Respon dalam waktu kurang dari 1000ms
  });

  // Test: Validasi respon waktu (response time) untuk FindByLimit endpoint
  test('Response time for FindByLimit Aplikasi', async ({ request }) => {
    const limit = 5;
    const response = await request.get(`${BASE_URL}?limit=${limit}`, {
      headers: {
        'Authorization': `Basic ${Buffer.from(`${USERNAME}:${PASSWORD}`).toString('base64')}`
      }
    });
    expect(response.status()).toBe(200);
    expect(response.time()).toBeLessThan(1000); // Respon dalam waktu kurang dari 1000ms
  });
});
