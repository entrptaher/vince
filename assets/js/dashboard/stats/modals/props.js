import React, { useCallback } from "react";

import Modal from './modal'
import { addFilter } from '../../query'
import { specialTitleWhenGoalFilter } from "../behaviours/goal-conversions";
import { EVENT_PROPS_PREFIX, hasGoalFilter } from "../../util/filters"
import BreakdownModal from "./breakdown-modal";
import * as metrics from "../reports/metrics";
import * as url from "../../util/url";
import { revenueAvailable } from "../../query";
import { useQueryContext } from "../../query-context";
import { useSiteContext } from "../../site-context";
import { customPropsRoute } from "../../router";

function PropsModal() {
  const { query } = useQueryContext();
  const site = useSiteContext();
  const { propKey } = customPropsRoute.useParams()

  /*global BUILD_EXTRA*/
  const showRevenueMetrics = BUILD_EXTRA && revenueAvailable(query, site)

  const reportInfo = {
    title: specialTitleWhenGoalFilter(query, 'Custom Property Breakdown'),
    dimension: propKey,
    endpoint: url.apiPath(site, `/custom-prop-values/${propKey}`),
    dimensionLabel: propKey
  }

  const getFilterInfo = useCallback((listItem) => {
    return {
      prefix: `${EVENT_PROPS_PREFIX}${propKey}`,
      filter: ["is", `${EVENT_PROPS_PREFIX}${propKey}`, [listItem.name]]
    }
  }, [propKey])

  const addSearchFilter = useCallback((query, searchString) => {
    return addFilter(query, ['contains', `${EVENT_PROPS_PREFIX}${propKey}`, [searchString]])
  }, [propKey])

  function chooseMetrics() {
    return [
      metrics.createVisitors({ renderLabel: (_query) => "Visitors" }),
      metrics.createEvents({ renderLabel: (_query) => "Events" }),
      hasGoalFilter(query) && metrics.createConversionRate(),
      !hasGoalFilter(query) && metrics.createPercentage(),
      showRevenueMetrics && metrics.createAverageRevenue(),
      showRevenueMetrics && metrics.createTotalRevenue(),
    ].filter(metric => !!metric)
  }

  return (
    <Modal>
      <BreakdownModal
        reportInfo={reportInfo}
        metrics={chooseMetrics()}
        getFilterInfo={getFilterInfo}
        addSearchFilter={addSearchFilter}
      />
    </Modal>
  )
}

export default PropsModal