/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package com.platform.quality.job.entity;

import javax.persistence.DiscriminatorValue;
import javax.persistence.Entity;

@Entity
@DiscriminatorValue("griffinStreamingJob")
public class StreamingJob extends AbstractJob {
    private static final long serialVersionUID = 3292253488392308505L;

    @Override
    public String getType() {
        return JobType.STREAMING.getName();
    }

    public StreamingJob() {
    }

    public StreamingJob(Long measureId, String jobName, String name,
                        String group, boolean deleted) {
        super(measureId, jobName, name, group, deleted);
        this.metricName = jobName;
    }
}
