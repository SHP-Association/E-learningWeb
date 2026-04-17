"""
Custom response utilities for standardizing API responses.

This module provides helper functions to ensure all API responses follow
a consistent format with proper data wrapping and metadata.
"""
import os
from rest_framework.response import Response
from rest_framework import status


def api_response(data=None, status_code=status.HTTP_200_OK, message=None, errors=None):
    """
    Standardized API response for single objects.
    
    Args:
        data: The response data (dict or serialized object)
        status_code: HTTP status code
        message: Optional message string
        errors: Optional error details
    
    Returns:
        Response object with standardized format:
        {
            "ci_environment": "",
            "data": { ... }
        }
    """
    response_data = {
        "ci_environment": os.getenv("CI_ENVIRONMENT", ""),
        "data": data if data is not None else {}
    }
    
    if message:
        response_data["message"] = message
    
    if errors:
        response_data["errors"] = errors
    
    return Response(response_data, status=status_code)


def api_list_response(items, total=None, status_code=status.HTTP_200_OK, message=None):
    """
    Standardized API response for list/collection endpoints.
    
    Args:
        items: List of items (already serialized)
        total: Total count of items (if None, will use len(items))
        status_code: HTTP status code
        message: Optional message string
    
    Returns:
        Response object with standardized format:
        {
            "ci_environment": "",
            "data": {
                "items": [...],
                "total": N
            }
        }
    """
    if total is None:
        total = len(items) if items else 0
    
    response_data = {
        "ci_environment": os.getenv("CI_ENVIRONMENT", ""),
        "data": {
            "items": items if items is not None else [],
            "total": total
        }
    }
    
    if message:
        response_data["message"] = message
    
    return Response(response_data, status=status_code)


def api_error_response(message, errors=None, status_code=status.HTTP_400_BAD_REQUEST):
    """
    Standardized API error response.
    
    Args:
        message: Error message string
        errors: Optional detailed error information
        status_code: HTTP status code
    
    Returns:
        Response object with standardized error format:
        {
            "ci_environment": "",
            "message": "Error message",
            "errors": { ... }
        }
    """
    response_data = {
        "ci_environment": os.getenv("CI_ENVIRONMENT", ""),
        "message": message,
        "success": False
    }
    
    if errors:
        response_data["errors"] = errors
    
    return Response(response_data, status=status_code)


def api_success_response(message, data=None, status_code=status.HTTP_200_OK):
    """
    Standardized API success response with message.
    
    Args:
        message: Success message string
        data: Optional response data
        status_code: HTTP status code
    
    Returns:
        Response object with standardized format
    """
    response_data = {
        "ci_environment": os.getenv("CI_ENVIRONMENT", ""),
        "message": message,
        "success": True
    }
    
    if data is not None:
        response_data["data"] = data
    
    return Response(response_data, status=status_code)
