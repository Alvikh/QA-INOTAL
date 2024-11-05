const { test, expect } = require('@playwright/test');

const baseURL = 'http://localhost:8080/sistem/v010/aplikasi';

// Uji endpoint GET /
test('GET / - Retrieve all aplikasi', async ({ request }) => {
  const response = await request.get(baseURL);
  expect(response.status()).toBe(200);
  const data = await response.json();
  expect(Array.isArray(data)).toBeTruthy(); // Pastikan respons berformat array
  if (data.length > 0) {
    expect(data[0]).toHaveProperty('kd');
    expect(data[0]).toHaveProperty('nama');
    expect(data[0]).toHaveProperty('label');
    expect(data[0]).toHaveProperty('logo');
    expect(data[0]).toHaveProperty('url_fe');
    expect(data[0]).toHaveProperty('url_api');
  }
});

// Uji endpoint GET /:kd
test('GET /:kd - Retrieve specific aplikasi', async ({ request }) => {
  const kd = 1; // Ganti dengan kode yang valid di database
  const response = await request.get(`${baseURL}/${kd}`);
  expect(response.status()).toBe(200);
  const data = await response.json();
  expect(data).toHaveProperty('kd');
  expect(data).toHaveProperty('nama');
  expect(data).toHaveProperty('label');
  expect(data).toHaveProperty('logo');
  expect(data).toHaveProperty('url_fe');
  expect(data).toHaveProperty('url_api');
});

// Uji endpoint POST /
test('POST / - Create new aplikasi', async ({ request }) => {
  const newAplikasi = {
    nama: 'Aplikasi Baru',
    label: 'Label Baru',
    logo: 'logo.png',
    url_fe: 'https://frontend.example.com',
    url_api: 'https://api.example.com'
  };
  const response = await request.post(baseURL, {
    data: newAplikasi,
  });
  expect(response.status()).toBe(201);
  const data = await response.json();
  expect(data).toHaveProperty('kd'); // Pastikan properti 'kd' ada pada respons
  expect(data.nama).toBe(newAplikasi.nama);
  expect(data.label).toBe(newAplikasi.label);
  expect(data.logo).toBe(newAplikasi.logo);
  expect(data.url_fe).toBe(newAplikasi.url_fe);
  expect(data.url_api).toBe(newAplikasi.url_api);
});

// Uji endpoint PUT /
test('PUT / - Update aplikasi', async ({ request }) => {
  const updatedAplikasi = {
    kd: 1, // Ganti dengan kode aplikasi yang ingin diperbarui
    nama: 'Aplikasi Diperbarui',
    label: 'Label Diperbarui',
    logo: 'logo_diperbarui.png',
    url_fe: 'https://frontend.updated.com',
    url_api: 'https://api.updated.com'
  };
  const response = await request.put(baseURL, {
    data: updatedAplikasi,
  });
  expect(response.status()).toBe(200);
  const data = await response.json();
  expect(data).toHaveProperty('status', 'success'); // Sesuaikan jika ada respons berbeda
  expect(data).toHaveProperty('kd', updatedAplikasi.kd);
  expect(data.nama).toBe(updatedAplikasi.nama);
  expect(data.label).toBe(updatedAplikasi.label);
  expect(data.logo).toBe(updatedAplikasi.logo);
  expect(data.url_fe).toBe(updatedAplikasi.url_fe);
  expect(data.url_api).toBe(updatedAplikasi.url_api);
});

// Uji endpoint DELETE /:kd
test('DELETE /:kd - Delete aplikasi', async ({ request }) => {
  const kd = 1; // Ganti dengan kode yang valid di database
  const response = await request.delete(`${baseURL}/${kd}`);
  expect(response.status()).toBe(204); // Biasanya 204 untuk penghapusan sukses
});
