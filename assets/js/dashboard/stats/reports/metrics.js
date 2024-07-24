import { hasGoalFilter } from "../../util/filters"
import numberFormatter, { durationFormatter, percentageFormatter } from "../../util/number-formatter"
import React from "react"

/*global BUILD_EXTRA*/
/*global require*/
function maybeRequire() {
  if (BUILD_EXTRA) {
    return require('../../extra/money')
  } else {
    return { default: null }
  }
}

const Money = maybeRequire().default

// Class representation of a metric.

// Metric instances can be created directly via the Metric constructor,
// or using special creator functions like `createVisitors`, which just
// fill out the known fields for that metric.

// ### Required props

// * `key` - the key under which to read values under in an API

// * `renderValue` - a function that takes a value of this metric, and
//   and returns the "rendered" version of it. Can be JSX or a string.

// * `renderLabel` - a function rendering a label for this metric given a
//   query argument. Can return JSX or string.

// ### Optional props

// * `meta` - a map with extra context for this metric. E.g. `plot`, or
//   `hiddenOnMobile` define some special behaviours in the context where
//   it's used.
export class Metric {
  constructor(props) {
    if (!props.key) {
      throw Error("Required field `key` is missing")
    }
    if (typeof props.renderLabel !== 'function') {
      throw Error("Required field `renderLabel` should be a function")
    }
    if (typeof props.renderValue !== 'function') {
      throw Error("Required field `renderValue` should be a function")
    }

    this.key = props.key
    this.renderValue = props.renderValue
    this.renderLabel = props.renderLabel
    this.meta = props.meta || {}
  }
}

// Creates a Metric class representing the `visitors` metric.

// Optional props for conveniently generating the `renderLabel` function:

// * `defaultLabel` - label when not realtime, and no goal filter applied
// * `realtimeLabel` - label when realtime period
// * `goalFilterLabel` - label when goal filter is applied
export const createVisitors = (props) => {
  let renderValue
  
  if (typeof props.renderValue === 'function') {
    renderValue = props.renderValue
  } else {
    renderValue = renderNumberWithTooltip
  }
  
  let renderLabel
  
  if (typeof props.renderLabel === 'function') {
    renderLabel = props.renderLabel
  } else {
    renderLabel = (query) => {
      const defaultLabel = props.defaultLabel || 'Visitors'
      const realtimeLabel = props.realtimeLabel || 'Current visitors'
      const goalFilterLabel = props.goalFilterLabel || 'Conversions'

      if (query.period === 'realtime') { return realtimeLabel }
      if (query && hasGoalFilter(query)) { return goalFilterLabel }
      return defaultLabel
    }
  }

  return new Metric({...props, key: "visitors", renderValue, renderLabel})
}

export const createConversionRate = (props) => {
  const renderValue = percentageFormatter
  const renderLabel = (_query) => "CR"
  return new Metric({...props, key: "conversion_rate", renderLabel, renderValue})
}

export const createPercentage = (props) => {
  const renderValue = (value) => value
  const renderLabel = (_query) => "%"
  return new Metric({...props, key: "percentage", renderLabel, renderValue})
}

export const createEvents = (props) => {
  const renderValue = typeof props.renderValue === 'function' ? props.renderValue : renderNumberWithTooltip
  return new Metric({...props, key: "events", renderValue: renderValue})
}

export const createTotalRevenue = (props) => {
  const renderValue = (value) => <Money formatted={value} />
  const renderLabel = (_query) => "Revenue"
  return new Metric({...props, key: "total_revenue", renderValue, renderLabel})
}

export const createAverageRevenue = (props) => {
  const renderValue = (value) => <Money formatted={value} />
  const renderLabel = (_query) => "Average"
  return new Metric({...props, key: "average_revenue", renderValue, renderLabel})
}

export const createTotalVisitors = (props) => {
  const renderValue = renderNumberWithTooltip
  const renderLabel = (_query) => "Total Visitors"
  return new Metric({...props, key: "total_visitors", renderValue, renderLabel})
}

export const createVisits = (props) => {
  const renderValue = renderNumberWithTooltip
  return new Metric({...props, key: "visits", renderValue})
}

export const createVisitDuration = (props) => {
  const renderValue = durationFormatter
  const renderLabel = (_query) => "Visit Duration"
  return new Metric({...props, key: "visit_duration", renderValue, renderLabel})
}

export const createBounceRate = (props) => {
  const renderValue = (value) => `${value}%`
  const renderLabel = (_query) => "Bounce Rate"
  return new Metric({...props, key: "bounce_rate", renderValue, renderLabel})
}

export const createPageviews = (props) => {
  const renderValue = renderNumberWithTooltip
  const renderLabel = (_query) => "Pageviews"
  return new Metric({...props, key: "pageviews", renderValue, renderLabel})
}

export const createTimeOnPage = (props) => {
  const renderValue = durationFormatter
  const renderLabel = (_query) => "Time on Page"
  return new Metric({...props, key: "time_on_page", renderValue, renderLabel})
}

export const createExitRate = (props) => {
  const renderValue = percentageFormatter
  const renderLabel = (_query) => "Exit Rate"
  return new Metric({...props, key: "exit_rate", renderValue, renderLabel})
}

export function renderNumberWithTooltip(value) {
  return <span tooltip={value}>{numberFormatter(value)}</span>
}