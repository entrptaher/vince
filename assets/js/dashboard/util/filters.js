import { useMemo } from "react"
import * as api from '../api'
import { useQueryContext } from '../query-context'

export const FILTER_MODAL_TO_FILTER_GROUP = {
  'page': ['page', 'entry_page', 'exit_page'],
  'source': ['source', 'referrer'],
  'location': ['country', 'region', 'city'],
  'screen': ['screen'],
  'browser': ['browser', 'browser_version'],
  'os': ['os', 'os_version'],
  'utm': ['utm_medium', 'utm_source', 'utm_campaign', 'utm_term', 'utm_content'],
  'goal': ['goal'],
  'props': ['props'],
  'hostname': ['hostname']
}

export const FILTER_GROUP_TO_MODAL_TYPE = Object.fromEntries(
  Object.entries(FILTER_MODAL_TO_FILTER_GROUP)
    .flatMap(([modalName, filterGroups]) => filterGroups.map((filterGroup) => [filterGroup, modalName]))
)

export const NO_CONTAINS_OPERATOR = new Set(['goal', 'screen'].concat(FILTER_MODAL_TO_FILTER_GROUP['location']))

export const EVENT_PROPS_PREFIX = "props:"

export const FILTER_OPERATIONS = {
  is: 'is',
  isNot: 'is_not',
  contains: 'contains',
  does_not_contain: 'does_not_contain'
};

export const FILTER_OPERATIONS_DISPLAY_NAMES = {
  [FILTER_OPERATIONS.is]: 'is',
  [FILTER_OPERATIONS.isNot]: 'is not',
  [FILTER_OPERATIONS.contains]: 'contains',
  [FILTER_OPERATIONS.does_not_contain]: 'does not contain'
}

const OPERATION_PREFIX = {
  [FILTER_OPERATIONS.isNot]: '!',
  [FILTER_OPERATIONS.contains]: '~',
  [FILTER_OPERATIONS.is]: ''
};


export function supportsIsNot(filterName) {
  return !['goal', 'prop_key'].includes(filterName)
}

export function isFreeChoiceFilter(filterName) {
  return !NO_CONTAINS_OPERATOR.has(filterName)
}

// As of March 2023, Safari does not support negative lookbehind regexes. In case it throws an error, falls back to plain | matching. This means
// escaping pipe characters in filters does not currently work in Safari
let NON_ESCAPED_PIPE_REGEX;
try {
  NON_ESCAPED_PIPE_REGEX = new RegExp("(?<!\\\\)\\|", "g")
} catch (_e) {
  NON_ESCAPED_PIPE_REGEX = '|'
}

const ESCAPED_PIPE = '\\|'

export function getLabel(labels, filterKey, value) {
  if (['country', 'region', 'city'].includes(filterKey)) {
    return labels[value]
  } else {
    return value
  }
}

export function getPropertyKeyFromFilterKey(filterKey) {
  return filterKey.slice(EVENT_PROPS_PREFIX.length)
}

export function getFiltersByKeyPrefix(query, prefix) {
  return query.filters.filter(([_operation, filterKey, _clauses]) => filterKey.startsWith(prefix))
}

function omitFiltersByKeyPrefix(query, prefix) {
  return query.filters.filter(([_operation, filterKey, _clauses]) => !filterKey.startsWith(prefix))
}

export function replaceFilterByPrefix(query, prefix, filter) {
  return omitFiltersByKeyPrefix(query, prefix).concat([filter])
}

export function isFilteringOnFixedValue(query, filterKey, expectedValue) {
  const filters = query.filters.filter(([_operation, key]) => filterKey == key)
  if (filters.length == 1) {
    const [operation, _filterKey, clauses] = filters[0]
    return operation === FILTER_OPERATIONS.is && clauses.length === 1 && (!expectedValue || clauses[0] == expectedValue)
  }
  return false
}

export function hasGoalFilter(query) {
  return getFiltersByKeyPrefix(query, "goal").length > 0
}

export function useHasGoalFilter() {
  const { query: { filters } } = useQueryContext();
  return useMemo(() => getFiltersByKeyPrefix({ filters }, "goal").length > 0, [filters]);
}

export function isRealTimeDashboard(query) {
  return query?.period === 'realtime'
}

export function useIsRealtimeDashboard() {
  const { query: { period } } = useQueryContext();
  return useMemo(() => isRealTimeDashboard({ period }), [period]);
}


// Note: Currently only a single goal filter can be applied at a time.
export function getGoalFilter(query) {
  return getFiltersByKeyPrefix(query, "goal")[0] || null
}

export function formatFilterGroup(filterGroup) {
  if (filterGroup === 'utm') {
    return 'UTM tags'
  } else if (filterGroup === 'location') {
    return 'Location'
  } else if (filterGroup === 'props') {
    return 'Property'
  } else {
    return formattedFilters[filterGroup]
  }
}

export function cleanLabels(filters, labels, mergedFilterKey, mergedLabels) {
  const filteredBy = Object.fromEntries(
    filters
      .flatMap(([_operation, filterKey, clauses]) => ['country', 'region', 'city'].includes(filterKey) ? clauses : [])
      .map((value) => [value, true])
  )
  let result = { ...labels }
  for (const value in labels) {
    if (!filteredBy[value]) {
      delete result[value]
    }
  }

  if (mergedFilterKey && ['country', 'region', 'city'].includes(mergedFilterKey)) {
    result = {
      ...result,
      ...mergedLabels
    }
  }

  return result
}

const EVENT_FILTER_KEYS = new Set(["name", "page", "goal", "hostname"])

export function serializeApiFilters(filters) {
  const apiFilters = filters.map(([operation, filterKey, clauses]) => {
    let apiFilterKey = `visit:${filterKey}`
    if (filterKey.startsWith(EVENT_PROPS_PREFIX) || EVENT_FILTER_KEYS.has(filterKey)) {
      apiFilterKey = `event:${filterKey}`
    }
    return [operation, apiFilterKey, clauses]
  })

  return JSON.stringify(apiFilters)
}

export function fetchSuggestions(apiPath, query, input, additionalFilter) {
  const updatedQuery = queryForSuggestions(query, additionalFilter)
  return api.get(apiPath, updatedQuery, { q: input.trim() })
}

function queryForSuggestions(query, additionalFilter) {
  let filters = query.filters
  if (additionalFilter) {
    const [_operation, filterKey, clauses] = additionalFilter

    // For suggestions, we remove already-applied filter with same key from query and add new filter (if feasible)
    if (clauses.length > 0) {
      filters = replaceFilterByPrefix(query, filterKey, additionalFilter)
    } else {
      filters = omitFiltersByKeyPrefix(query, filterKey)
    }
  }
  return { ...query, filters }
}

export function getFilterGroup([_operation, filterKey, _clauses]) {
  return filterKey.startsWith(EVENT_PROPS_PREFIX) ? 'props' : filterKey
}


export const formattedFilters = {
  'goal': 'Goal',
  'props': 'Property',
  'prop_key': 'Property',
  'prop_value': 'Value',
  'source': 'Source',
  'utm_medium': 'UTM Medium',
  'utm_source': 'UTM Source',
  'utm_campaign': 'UTM Campaign',
  'utm_content': 'UTM Content',
  'utm_term': 'UTM Term',
  'referrer': 'Referrer URL',
  'screen': 'Screen size',
  'browser': 'Browser',
  'browser_version': 'Browser Version',
  'os': 'Operating System',
  'os_version': 'Operating System Version',
  'country': 'Country',
  'region': 'Region',
  'city': 'City',
  'page': 'Page',
  'hostname': 'Hostname',
  'entry_page': 'Entry Page',
  'exit_page': 'Exit Page',
}


export function parseLegacyFilter(filterKey, rawValue) {
  const operation = Object.keys(OPERATION_PREFIX)
    .find(operation => OPERATION_PREFIX[operation] === rawValue[0]) || FILTER_OPERATIONS.is;

  const value = operation === FILTER_OPERATIONS.is ? rawValue : rawValue.substring(1)

  const clauses = value
    .split(NON_ESCAPED_PIPE_REGEX)
    .filter((clause) => !!clause)
    .map((val) => val.replaceAll(ESCAPED_PIPE, '|'))

  return [operation, filterKey, clauses]
}

export function parseLegacyPropsFilter(rawValue) {
  return Object.entries(JSON.parse(rawValue)).map(([key, propVal]) => {
    return parseLegacyFilter(`${EVENT_PROPS_PREFIX}${key}`, propVal)
  })
}
