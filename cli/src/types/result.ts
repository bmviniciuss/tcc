export type Result = {
  metadata: {
    type: string
    testConfig: {
      id: string
      name: string
      vus: number
      duration: string
      executedAt: string
    }
  }
  root_group: {
    name: string
    path: string
    id: string
    groups: Array<any>
    checks: Array<{
      path: string
      id: string
      passes: number
      fails: number
      name: string
    }>
  }
  options: {
    summaryTrendStats: Array<string>
    summaryTimeUnit: string
    noColor: boolean
  }
  state: {
    isStdOutTTY: boolean
    isStdErrTTY: boolean
    testRunDurationMs: number
  }
  metrics: {
    http_req_failed: {
      type: string
      contains: string
      values: {
        rate: number
        passes: number
        fails: number
      }
    }
    http_req_blocked: {
      type: string
      contains: string
      values: {
        min: number
        med: number
        avg: number
        max: number
        'p(90)': number
        'p(95)': number
        'p(99)': number
      }
    }
    data_received: {
      type: string
      contains: string
      values: {
        count: number
        rate: number
      }
    }
    vus_max: {
      type: string
      contains: string
      values: {
        max: number
        value: number
        min: number
      }
    }
    http_req_connecting: {
      type: string
      contains: string
      values: {
        med: number
        avg: number
        max: number
        'p(90)': number
        'p(95)': number
        'p(99)': number
        min: number
      }
    }
    checks: {
      type: string
      contains: string
      values: {
        fails: number
        rate: number
        passes: number
      }
      thresholds: {
        'rate>0.99': {
          ok: boolean
        }
      }
    }
    http_reqs: {
      type: string
      contains: string
      values: {
        count: number
        rate: number
      }
    }
    http_req_tls_handshaking: {
      type: string
      contains: string
      values: {
        'p(95)': number
        'p(99)': number
        min: number
        med: number
        avg: number
        max: number
        'p(90)': number
      }
    }
    http_req_sending: {
      type: string
      contains: string
      values: {
        min: number
        med: number
        avg: number
        max: number
        'p(90)': number
        'p(95)': number
        'p(99)': number
      }
    }
    vus: {
      type: string
      contains: string
      values: {
        value: number
        min: number
        max: number
      }
    }
    'http_req_duration{expected_response:true}': {
      type: string
      contains: string
      values: {
        min: number
        med: number
        avg: number
        max: number
        'p(90)': number
        'p(95)': number
        'p(99)': number
      }
    }
    http_req_receiving: {
      type: string
      contains: string
      values: {
        'p(95)': number
        'p(99)': number
        min: number
        med: number
        avg: number
        max: number
        'p(90)': number
      }
    }
    data_sent: {
      type: string
      contains: string
      values: {
        rate: number
        count: number
      }
    }
    http_req_duration: {
      type: string
      contains: string
      values: {
        min: number
        med: number
        avg: number
        max: number
        'p(90)': number
        'p(95)': number
        'p(99)': number
      }
    }
    iteration_duration: {
      values: {
        med: number
        avg: number
        max: number
        'p(90)': number
        'p(95)': number
        'p(99)': number
        min: number
      }
      type: string
      contains: string
    }
    iterations: {
      type: string
      contains: string
      values: {
        count: number
        rate: number
      }
    }
    http_req_waiting: {
      type: string
      contains: string
      values: {
        'p(90)': number
        'p(95)': number
        'p(99)': number
        min: number
        med: number
        avg: number
        max: number
      }
    }
  }
}
