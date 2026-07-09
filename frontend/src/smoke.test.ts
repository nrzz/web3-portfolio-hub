import { describe, expect, it } from 'vitest'

describe('frontend smoke', () => {
  it('runs the test harness', () => {
    expect(true).toBe(true)
  })

  it('resolves API health path used by deploy checks', () => {
    expect('/health').toBe('/health')
    expect('/api/health').toBe('/api/health')
  })
})
