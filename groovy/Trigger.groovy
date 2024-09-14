// Import necessary Jenkins classes
import hudson.model.*
import jenkins.model.*

println "Environment Variables:"
System.getenv().each { key, value ->
    println "${key} = ${value}"
}

// Execute 'ls' command to list directory contents
println "Directory Contents:"
def proc = "ls".execute()
proc.in.eachLine { line ->
    println line
}
proc.waitFor()

// Get the environment variable 'Alpha'
def alpha = System.getenv('Alpha')

// Check if 'Alpha' is true
if (alpha == 'true') {
    // Get the Jenkins instance
    def jenkins = Jenkins.getInstance()

    // Specify the job name to trigger
    def jobName = "test-pipeline"

    // Get the job by name
    def job = jenkins.getItemByFullName(jobName)

    if (job) {
        // Trigger the job
        job.scheduleBuild(new Cause.UserIdCause())
        println "Triggered build for job: ${jobName}"
    } else {
        println "Job not found: ${jobName}"
    }
} else {
    println "Environment variable 'Alpha' is not set to true. No build triggered."
}
