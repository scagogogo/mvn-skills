// Headless Mermaid syntax validator for the docs.
// Extracts every ```mermaid fenced block from the given markdown files and runs
// mermaid.parse() on each under jsdom, so syntax errors are caught at build/CI
// time instead of silently producing a broken diagram in the browser.
import { readFileSync } from 'node:fs'
import { JSDOM } from 'jsdom'

const dom = new JSDOM('<!DOCTYPE html><body></body>', { pretendToBeVisual: true })
globalThis.window = dom.window
globalThis.document = dom.window.document
if (!globalThis.navigator) {
  Object.defineProperty(globalThis, 'navigator', { value: dom.window.navigator, configurable: true })
}

const { default: mermaid } = await import('mermaid')
mermaid.initialize({ startOnLoad: false, securityLevel: 'loose' })

const files = process.argv.slice(2)
if (files.length === 0) {
  console.error('usage: node validate-mermaid.mjs <file.md> [...]')
  process.exit(2)
}

let total = 0
let failed = 0

for (const file of files) {
  const src = readFileSync(file, 'utf8')
  const re = /```mermaid\n([\s\S]*?)```/g
  let m
  let idx = 0
  while ((m = re.exec(src)) !== null) {
    idx++
    total++
    const graph = m[1]
    try {
      await mermaid.parse(graph)
      console.log(`  ok   ${file} #${idx}`)
    } catch (err) {
      failed++
      const firstLine = graph.trim().split('\n')[0]
      console.error(`  FAIL ${file} #${idx} (${firstLine})`)
      console.error(`       ${String(err.message || err).split('\n')[0]}`)
    }
  }
}

console.log(`\n${total - failed}/${total} diagrams valid`)
process.exit(failed === 0 ? 0 : 1)
