
// Function to toggle between sections of the page
function showHomePage() {
  // Show the home section and hide other sections
  document.getElementById("homeSection").style.display = "block";
  document.getElementById("courseListSection").style.display = "none";
  document.getElementById("courseDetailSection").style.display = "none";
}

function showCoursePage() {
  // Show the course list section and hide other sections
  document.getElementById("homeSection").style.display = "none";
  document.getElementById("courseListSection").style.display = "block";
  document.getElementById("courseDetailSection").style.display = "none";
}

function showCourseDetail(courseTitle, courseDescription) {
  // Display the course details by changing the text and showing the detail section
  document.getElementById("course-title").innerText = courseTitle;
  document.getElementById("course-description").innerText = courseDescription;
  document.getElementById("homeSection").style.display = "none";
  document.getElementById("courseListSection").style.display = "none";
  document.getElementById("courseDetailSection").style.display = "block";
}

// Populate course list on the homepage
function populateCourses() {
  const courseList = document.querySelector(".course-list"); // Select the course list container

  // Loop through the predefined courses and create course items
  courses.forEach((course) => {
    const courseItem = document.createElement("div"); // Create a new div element for each course
    courseItem.className = "course-item"; // Add a class to the course item for styling

    // Set the inner HTML with course details and a link to view details
    courseItem.innerHTML = `
            <h3>${course.title}</h3>
            <p>${course.description}</p>
            <a href="#" onclick="showCourseDetail('${course.title}', '${course.description}')">Learn More</a>
        `;
    courseList.appendChild(courseItem); // Append the course item to the course list
  });
}

// Initialize the homepage
if (document.querySelector(".course-list")) {
  populateCourses(); // Populate course list on the homepage
}

// Handle modal functionality for enrollment
document.addEventListener("DOMContentLoaded", () => {
  const enrollBtn = document.getElementById("enroll-btn");
  const modal = document.getElementById("enrollModal");
  const closeBtn = document.querySelector(".close-btn");

  // Show the modal when "Enroll" is clicked
  enrollBtn.addEventListener("click", () => {
    modal.style.display = "block";
  });

  // Hide the modal when the close button is clicked
  closeBtn.addEventListener("click", () => {
    modal.style.display = "none";
  });

  // Hide the modal if clicking outside of it
  window.addEventListener("click", (e) => {
    if (e.target == modal) {
      modal.style.display = "none";
    }
  });
});

// Function to handle course enrollment (Google Form)
function enroll() {
  const googleFormUrl = "https://forms.gle/mCEVQo8qzRf6748J8"; // Replace with actual form URL
  window.open(googleFormUrl, "_blank"); // Open the form in a new tab
}